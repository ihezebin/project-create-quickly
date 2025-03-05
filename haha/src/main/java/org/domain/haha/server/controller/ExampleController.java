package org.domain.haha.server.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import org.domain.haha.application.service.ExampleService;
import org.domain.haha.config.CustomizeConfig;
import org.domain.haha.domain.entity.Example;
import org.domain.haha.application.dto.ResponseBody;
import org.domain.haha.application.dto.example.InsertExampleRequest;
import org.domain.haha.application.dto.example.FindExampleRequest;
import jakarta.validation.Valid;

@RestController
@RequestMapping("/example")
public class ExampleController {

    @Autowired
    private ExampleService exampleService;

    @Autowired
    private CustomizeConfig customizeConfig;

    @PostMapping("/insert")
    public ResponseBody<Example> insert(@Valid @RequestBody InsertExampleRequest request) {
        Example example = new Example();
        example.setUsername(request.getUsername());
        example.setPassword(request.getPassword());
        example.setEmail(request.getEmail());

        example = exampleService.insert(example);
        return ResponseBody.success(example);
    }

    @GetMapping("/find")
    public ResponseBody<Example> findOne(@Valid @ModelAttribute FindExampleRequest request) {
        String username = request.getUsername();
        Example example = exampleService.findOne(username);
        return ResponseBody.success(example);
    }

    @GetMapping("/find_es")
    public ResponseBody<Example> findEs(@Valid @ModelAttribute FindExampleRequest request) {
        String username = request.getUsername();
        Example example = exampleService.findEs(username);
        return ResponseBody.success(example);
    }

    @GetMapping("/find_mongo")
    public ResponseBody<Example> findMongo(@Valid @ModelAttribute FindExampleRequest request) {
        String username = request.getUsername();
        Example example = exampleService.findMongo(username);
        return ResponseBody.success(example);
    }

    @GetMapping("/find_redis")
    public ResponseBody<Example> findRedis(@Valid @ModelAttribute FindExampleRequest request) {
        String username = request.getUsername();
        Example example = exampleService.findRedis(username);
        return ResponseBody.success(example);
    }

    @GetMapping("/config")
    public CustomizeConfig config() {
        return customizeConfig;
    }
}
