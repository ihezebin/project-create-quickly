package org.domain.haha.config;

import org.apache.http.HttpHost;
import org.elasticsearch.client.RestClient;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;

import co.elastic.clients.elasticsearch.ElasticsearchClient;
import co.elastic.clients.json.jackson.JacksonJsonpMapper;
import co.elastic.clients.transport.rest_client.RestClientTransport;
import co.elastic.clients.transport.ElasticsearchTransport;

import java.util.Arrays;

@Configuration
@ConditionalOnProperty(prefix = "elasticsearch", name = "uris", matchIfMissing = false)
public class ElasticsearchConfig {

        @Value("${elasticsearch.uris}")
        private String[] uris;

        @Bean
        public ElasticsearchClient elasticsearchClient() {
                // Create the low-level client
                RestClient restClient = RestClient.builder(
                                Arrays.stream(uris)
                                                .map(HttpHost::create)
                                                .toArray(HttpHost[]::new))
                                .build();

                // Create the transport with a Jackson mapper
                ElasticsearchTransport transport = new RestClientTransport(
                                restClient,
                                new JacksonJsonpMapper());

                // And create the API client
                return new ElasticsearchClient(transport);
        }
}