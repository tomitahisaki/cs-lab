class IdempotentDemoJob < ApplicationJob
  queue_as :default

  # 引数:
  #   token: 任意の識別子（本来は注文IDなど）。同じtokenなら一度だけ実行される。
  #   fail_once: trueにすると最初の1回だけ失敗し、Sidekiqの自動リトライで成功に至る挙動を確認できる。
  def perform(token, fail_once: false, work_ms: 200)
    done_key      = "demo:#{token}:done"        # 完了フラグ
    inflight_key  = "demo:#{token}:processing"  # 同時実行抑制（ロック）
    attempts_key  = "demo:#{token}:attempts"    # 失敗を1回だけ起こすカウンタ

    # すでに完了していれば何もしない（冪等）
    if redis_get(done_key)
      Rails.logger.info "✅ already done token=#{token}"
      return
    end

    # 最初の1回だけ失敗させる（リトライの挙動を観察）
    if fail_once && redis_incr(attempts_key) == 1
      raise "temporary failure for token=#{token}"
    end

    # 同時実行を抑制（5分のロック）
    locked = redis_set_nx(inflight_key, 1, ex: 300)
    unless locked
      Rails.logger.info "⏭️ skip: already processing token=#{token}"
      return
    end

    begin
      # 二重実行ガード（リトライ等で再突入しても安全）
      if redis_get(done_key)
        Rails.logger.info "✅ already done token=#{token}"
        return
      end

      # 擬似的な仕事
      sleep(work_ms.to_f / 1000.0)
      Rails.logger.info "🛠️ do work token=#{token}"

      # 完了マーク（以後は即returnされる）
      redis_set(done_key, Time.now.to_i)
    ensure
      # ロック解放
      redis_del(inflight_key)
    end
  end

  private

  # --- Redis helpers ---
  def sidekiq_redis(&blk)
    # Sidekiqの接続をそのまま使う（別gemは不要）
    Sidekiq.redis(&blk)
  end

  def redis_get(key)
    sidekiq_redis { |r| r.get(key) }
  end

  def redis_set(key, value, **opts)
    sidekiq_redis { |r| r.set(key, value, **opts) }
  end

  def redis_set_nx(key, value, ex: nil)
    sidekiq_redis { |r| r.set(key, value, nx: true, ex: ex) }
  end

  def redis_incr(key)
    sidekiq_redis { |r| r.incr(key) }
  end

  def redis_del(key)
    sidekiq_redis { |r| r.del(key) }
  end
end
