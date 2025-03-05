package org.domain.haha.cron;

import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Component
@ConditionalOnProperty(prefix = "cron", name = "enabled", havingValue = "true")
/*
 * * * * * * ?
 * │ │ │ │ │ │
 * │ │ │ │ │ └ 星期（1-7）
 * │ │ │ │ └── 月（1-12）
 * │ │ │ └──── 日（1-31）
 * │ │ └────── 时（0-23）
 * │ └──────── 分（0-59）
 * └────────── 秒（0-59）
 */
public class ExampleCron {

    // 每5秒执行一次
    @Scheduled(fixedRate = 5000)
    public void task1() {
        log.info("定时任务1执行");
    }

    // 每天凌晨1点执行
    @Scheduled(cron = "0 0 1 * * ?")
    public void task2() {
        log.info("定时任务2执行");
    }

    // 每分钟的第30秒执行
    @Scheduled(cron = "30 * * * * ?")
    public void task3() {
        try {
            log.info("定时任务3执行");
        } catch (Exception e) {
            log.error("定时任务3执行失败: ", e);
        }
    }
}