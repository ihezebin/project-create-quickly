package org.domain.haha.domain.repository;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;
import org.springframework.stereotype.Repository;
import org.domain.haha.domain.entity.Example;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;

@Repository("exampleMapper")
@Mapper
@ConditionalOnProperty(prefix = "spring.datasource", name = "url", matchIfMissing = false)
public interface ExampleMapper extends ExampleRepository {

    @Insert("INSERT INTO example (id, username, password, email, salt) VALUES (#{id}, #{username}, #{password}, #{email}, #{salt})")
    public void insertOne(Example example);

    @Select("SELECT * FROM example WHERE username = #{username}")
    public Example findByUsername(String username);

    @Select("SELECT * FROM example WHERE email = #{email}")
    public Example findByEmail(String email);

    @Select("SELECT COUNT(*) FROM example")
    long count();

    @Select("SELECT * FROM example LIMIT #{offset}, #{pageSize}")
    List<Example> findPage(@Param("offset") int offset, @Param("pageSize") int pageSize);
}
