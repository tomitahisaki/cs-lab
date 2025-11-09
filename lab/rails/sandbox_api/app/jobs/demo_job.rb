class DemoJob < ApplicationJob
  queue_as :default

  def perform(name = "CS-lab")
    Rails.logger.info("[DemoJob] Hello #{name}! pid=#{Process.pid} tid=#{Thread.current.object_id}")
    sleep 1
    Rails.logger.info("[DemoJob] Finished #{name}")
  end
end
