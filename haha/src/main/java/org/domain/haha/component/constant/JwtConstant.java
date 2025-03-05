package org.domain.haha.component.constant;

import io.jsonwebtoken.security.Keys;
import java.security.Key;

public class JwtConstant {
    // JWT密钥，实际项目中应该配置在配置文件中
    public static final Key SECRET_KEY = Keys.secretKeyFor(io.jsonwebtoken.SignatureAlgorithm.HS256);

    // Token过期时间（毫秒）
    public static final long EXPIRATION_TIME = 3600000; // 1小时

    // Token前缀
    public static final String TOKEN_PREFIX = "Bearer ";

    // Header中的Token键名
    public static final String HEADER_STRING = "Authorization";
}