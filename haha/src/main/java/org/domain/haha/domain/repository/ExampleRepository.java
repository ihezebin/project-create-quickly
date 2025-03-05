package org.domain.haha.domain.repository;

import org.domain.haha.domain.entity.Example;

public interface ExampleRepository {
    public void insertOne(Example example);

    public Example findByUsername(String username);

    public Example findByEmail(String email);
}
