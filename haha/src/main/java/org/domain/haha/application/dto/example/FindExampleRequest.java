package org.domain.haha.application.dto.example;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

@Data
public class FindExampleRequest {
    @NotBlank(message = "用户名不能为空")
    private String username;
}
