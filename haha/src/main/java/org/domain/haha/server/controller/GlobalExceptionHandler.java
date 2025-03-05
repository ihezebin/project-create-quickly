package org.domain.haha.server.controller;

import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.context.support.DefaultMessageSourceResolvable;
import java.util.stream.Collectors;

import org.domain.haha.application.dto.ResponseBody;
import org.domain.haha.application.dto.ResponseBodyCode;
import org.domain.haha.exception.ErrorException;

import lombok.extern.slf4j.Slf4j;

@RestControllerAdvice
@Slf4j
public class GlobalExceptionHandler {

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseBody<?> handleValidationException(MethodArgumentNotValidException ex) {
        String errorMsg = ex.getBindingResult().getAllErrors()
                .stream()
                .map(DefaultMessageSourceResolvable::getDefaultMessage)
                .collect(Collectors.joining(", "));
        return ResponseBody.error(ResponseBodyCode.BAD_REQUEST, errorMsg);
    }

    // 处理业务异常
    @ExceptionHandler(ErrorException.class)
    public ResponseBody<?> handleErrorException(ErrorException ex) {
        return ResponseBody.error(ex.getCode(), ex.getMessage());
    }

    // 其他异常
    @ExceptionHandler(Exception.class)
    public ResponseBody<?> handleUnexpectedException(Exception ex) {
        log.error("Unexpected exception: {}", ex.getMessage());

        return ResponseBody.error(ResponseBodyCode.INTERNAL_SERVER_ERROR);
    }
}