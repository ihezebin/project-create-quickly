package org.domain.haha.component.httpclient;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpHeaders;

import org.domain.haha.application.dto.ResponseBody;
import org.domain.haha.domain.entity.Example;
import org.domain.haha.exception.ErrorException;
import org.domain.haha.application.dto.ResponseBodyCode;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@Component
public class ExampleHttpClient {

    @Value("${example.service.base-url}")
    private String baseUrl;

    @Autowired
    private RestTemplate restTemplate;

    public Example getExampleById(String id) throws ErrorException {
        try {
            ResponseBody response = restTemplate.getForObject(
                    baseUrl + "/api/example/{id}",
                    ResponseBody.class,
                    id);

            if (response.getCode() != ResponseBodyCode.OK.getCode()) {
                throw new ErrorException(ResponseBodyCode.BAD_REQUEST, response.getMessage());
            }
            return (Example) response.getData();

        } catch (Exception e) {
            log.error("Failed to get example by id: {}", id, e);
            throw new ErrorException(ResponseBodyCode.INTERNAL_SERVER_ERROR, e);
        }
    }

    public Example createExample(Example example) throws ErrorException {
        try {
            ResponseBody response = restTemplate.postForObject(
                    baseUrl + "/api/example",
                    example,
                    ResponseBody.class);

            if (response.getCode() != ResponseBodyCode.OK.getCode()) {
                throw new ErrorException(ResponseBodyCode.BAD_REQUEST, response.getMessage());
            }
            return (Example) response.getData();

        } catch (Exception e) {
            log.error("Failed to create example: {}", example, e);
            throw new ErrorException(ResponseBodyCode.INTERNAL_SERVER_ERROR, e);
        }
    }

    public String getBaiduHomePage() throws ErrorException {
        try {
            ResponseEntity<String> response = restTemplate.getForEntity("http://www.baidu.com", String.class);

            // 处理重定向
            if (response.getStatusCode().is3xxRedirection()) {
                String redirectUrl = response.getHeaders().getFirst(HttpHeaders.LOCATION);
                return restTemplate.getForObject(redirectUrl, String.class);
            }

            return response.getBody();

        } catch (Exception e) {
            log.error("Failed to get Baidu homepage", e);
            throw new ErrorException(ResponseBodyCode.INTERNAL_SERVER_ERROR, e);
        }
    }
}