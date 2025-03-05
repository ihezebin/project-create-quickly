package org.domain.haha.worker;

import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Component;
import lombok.extern.slf4j.Slf4j;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;

@Slf4j
@Component
@ConditionalOnProperty(prefix = "worker", name = "enabled", havingValue = "true")
public class ExampleWorker {

    private volatile boolean running = true;
    private Thread workerThread;

    @PostConstruct
    public void start() {
        workerThread = new Thread(() -> {
            while (running) {
                try {
                    log.info("hello example");
                    Thread.sleep(5000); // 休眠5秒
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    break;
                } catch (Exception e) {
                    log.error("ExampleWorker执行失败: ", e);
                }
            }
        }, "example-worker");

        workerThread.start();
        log.info("ExampleWorker started");
    }

    @PreDestroy
    public void stop() {
        running = false;
        if (workerThread != null) {
            workerThread.interrupt();
            try {
                workerThread.join(60000); // 等待最多60秒让线程结束
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
        log.info("ExampleWorker stopped");
    }
}