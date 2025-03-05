package org.domain.haha;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

public class JwtTest {

    // 使用强密钥，实际项目中应该配置在配置文件中
    private static final Key key = Keys.secretKeyFor(SignatureAlgorithm.HS256);

    // mvn test -Dtest=org.domain.haha.JwtTest#testCreateAndParseToken
    @Test
    public void testCreateAndParseToken() {
        // 准备测试数据
        String username = "test_user";
        String role = "admin";

        // 创建 claims
        Map<String, Object> claims = new HashMap<>();
        claims.put("username", username);
        claims.put("role", role);

        // 获取当前时间
        Date now = new Date();
        // token 过期时间设置为1小时
        Date expiration = new Date(now.getTime() + 3600000);

        // 创建 token
        String token = Jwts.builder()
                .setClaims(claims)
                .setIssuedAt(now)
                .setExpiration(expiration)
                .signWith(key)
                .compact();

        // 验证 token 不为空
        assertNotNull(token);
        System.out.println("Generated Token: " + token);

        // 解析 token
        Claims parsedClaims = Jwts.parserBuilder()
                .setSigningKey(key)
                .build()
                .parseClaimsJws(token)
                .getBody();

        // 验证解析结果
        assertEquals(username, parsedClaims.get("username"));
        assertEquals(role, parsedClaims.get("role"));

        // 验证过期时间
        assertTrue(parsedClaims.getExpiration().after(now));
    }

    @Test
    public void testTokenExpiration() {
        // 创建一个已过期的 token
        Date now = new Date();
        Date expiration = new Date(now.getTime() - 1000); // 过期时间设置为1秒前

        String token = Jwts.builder()
                .setSubject("test_user")
                .setIssuedAt(now)
                .setExpiration(expiration)
                .signWith(key)
                .compact();

        // 验证解析过期 token 会抛出异常
        assertThrows(io.jsonwebtoken.ExpiredJwtException.class, () -> {
            Jwts.parserBuilder()
                    .setSigningKey(key)
                    .build()
                    .parseClaimsJws(token);
        });
    }

    @Test
    public void testInvalidSignature() {
        // 使用不同的密钥
        Key anotherKey = Keys.secretKeyFor(SignatureAlgorithm.HS256);

        String token = Jwts.builder()
                .setSubject("test_user")
                .signWith(key)
                .compact();

        // 验证使用错误的密钥解析会抛出异常
        assertThrows(io.jsonwebtoken.security.SecurityException.class, () -> {
            Jwts.parserBuilder()
                    .setSigningKey(anotherKey)
                    .build()
                    .parseClaimsJws(token);
        });
    }
}