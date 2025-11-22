class FlakyJob < ApplicationJob
  queue_as :default

  # FalkyJob.perform_later(1) => 失敗
  # FalkyJob.perform_later(2) => 成功
  def perform(i)
    raise "boom #{i}" if i.to_i.odd? # 奇数ならわざと失敗
    Rails.logger.info("✅ success🍏🍏 #{i}")
  end
end
