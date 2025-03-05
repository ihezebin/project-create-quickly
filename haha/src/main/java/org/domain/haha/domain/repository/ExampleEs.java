package org.domain.haha.domain.repository;

import org.springframework.stereotype.Repository;
import co.elastic.clients.elasticsearch.ElasticsearchClient;
import co.elastic.clients.elasticsearch.core.SearchResponse;
import co.elastic.clients.elasticsearch.core.search.Hit;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.autoconfigure.condition.ConditionalOnBean;

import org.domain.haha.application.dto.ResponseBodyCode;
import org.domain.haha.domain.entity.Example;
import org.domain.haha.exception.ErrorException;

@Slf4j
@Repository
@ConditionalOnBean(ElasticsearchClient.class)
public class ExampleEs implements ExampleRepository {

    private final ElasticsearchClient esClient;
    private static final String INDEX_NAME = "example";

    public ExampleEs(ElasticsearchClient esClient) {
        this.esClient = esClient;
    }

    public void insertOne(Example example) throws ErrorException {
        try {
            // PUT /example/_doc/{id}
            esClient.index(i -> i
                    .index(INDEX_NAME)
                    .id(example.getId())
                    .document(example));
        } catch (Exception e) {
            throw new ErrorException(ResponseBodyCode.INTERNAL_SERVER_ERROR, e);
        }
    }

    public Example findByUsername(String username) throws ErrorException {
        try {
            // GET /example/_search
            // {
            // "query": {
            // "match": {
            // "username": "xxx"
            // }
            // }
            // }
            SearchResponse<Example> response = esClient.search(s -> s
                    .index(INDEX_NAME)
                    .query(q -> q
                            .match(m -> m
                                    .field("username")
                                    .query(username))),
                    Example.class);

            return extractFirstHit(response);
        } catch (Exception e) {
            throw new ErrorException(ResponseBodyCode.INTERNAL_SERVER_ERROR, e);
        }
    }

    public Example findByEmail(String email) throws ErrorException {
        try {
            // GET /example/_search
            // {
            // "query": {
            // "term": {
            // "email": "xxx@xxx.com"
            // }
            // }
            // }
            SearchResponse<Example> response = esClient.search(s -> s
                    .index(INDEX_NAME)
                    .query(q -> q
                            .term(t -> t
                                    .field("email")
                                    .value(email))),
                    Example.class);

            return extractFirstHit(response);
        } catch (Exception e) {
            throw new ErrorException(ResponseBodyCode.INTERNAL_SERVER_ERROR, e);
        }
    }

    private Example extractFirstHit(SearchResponse<Example> response) {
        if (response.hits().hits().isEmpty()) {
            return null;
        }
        Hit<Example> hit = response.hits().hits().get(0);
        return hit.source();
    }
}
