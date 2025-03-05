package org.domain.haha;

import org.domain.haha.component.httpclient.ExampleHttpClient;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
public class ExampleHttpClientTest {

        @Autowired
        private ExampleHttpClient exampleHttpClient;

        // mvn test
        // -Dtest=org.domain.haha.ExampleHttpClientTest#getBaiduHomePage_Success
        @Test
        void getBaiduHomePage_Success() {
                // 执行测试
                String result = exampleHttpClient.getBaiduHomePage();

                // 打印结果
                System.out.println(result);

                // 验证结果
                assertNotNull(result);
                // 由于百度首页可能会有变化，我们放宽验证条件
                assertTrue(result.contains("百度") || result.contains("baidu"));
        }
}