package org.domain.haha.server.filter;

import jakarta.servlet.*;
import jakarta.servlet.annotation.WebFilter;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.annotation.Order;
import org.springframework.web.util.ContentCachingRequestWrapper;
import org.springframework.web.util.ContentCachingResponseWrapper;

import java.io.IOException;

@WebFilter(filterName = "requestLoggingFilter", urlPatterns = "/*")
@Order(2)
@Slf4j
public class RequestLoggingFilter implements Filter {

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
            throws IOException, ServletException {
        ContentCachingRequestWrapper requestWrapper = new ContentCachingRequestWrapper((HttpServletRequest) request);
        ContentCachingResponseWrapper responseWrapper = new ContentCachingResponseWrapper(
                (HttpServletResponse) response);

        long startTime = System.currentTimeMillis();

        try {
            // 记录请求信息
            logRequest(requestWrapper);

            chain.doFilter(requestWrapper, responseWrapper);

            // 记录响应信息
            logResponse(responseWrapper, System.currentTimeMillis() - startTime);
        } finally {
            // 复制响应内容到原始响应
            responseWrapper.copyBodyToResponse();
        }
    }

    private void logRequest(ContentCachingRequestWrapper request) throws IOException {
        String method = request.getMethod();
        String uri = request.getRequestURI();
        String queryString = request.getQueryString();
        if (queryString != null) {
            uri += "?" + queryString;
        }

        String body = new String(request.getContentAsByteArray());
        if (!body.isEmpty()) {
            log.info("request method: {}, uri: {} - body: {}", method, uri, body);
        } else {
            log.info("request method: {}, uri: {}", method, uri);
        }
    }

    private void logResponse(ContentCachingResponseWrapper response, long duration) throws IOException {
        String body = new String(response.getContentAsByteArray());
        log.info("response status: {}, time: {}ms, body: {}",
                response.getStatus(),
                duration,
                body.isEmpty() ? "null" : body);
    }
}