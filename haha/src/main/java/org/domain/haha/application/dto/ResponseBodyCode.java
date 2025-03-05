package org.domain.haha.application.dto;

import lombok.Getter;
import lombok.AllArgsConstructor;

@Getter
@AllArgsConstructor
public enum ResponseBodyCode {
    OK(0, "OK"),
    VALIDATE_RULE_FAILED(1, "Validate Rule Failed"),
    INTERNAL_SERVER_ERROR(2, "Internal Server Error"),
    BAD_REQUEST(3, "Bad Request"),
    UNAUTHORIZED(4, "Unauthorized"),
    NOT_FOUND(5, "Not Found"),
    FORBIDDEN(6, "Forbidden"),
    TIMEOUT(7, "Timeout"),
    CREATED(8, "Created"),
    ACCEPTED(9, "Accepted"),
    NO_CONTENT(10, "No Content"),
    RESET_CONTENT(11, "Reset Content"),
    AUTHORIZATION_FAILED(12, "Authorization Failed");

    private final int code;
    private final String message;
}
