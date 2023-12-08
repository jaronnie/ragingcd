package com.jaronnie.fronted_backend_admin.backend_java.controller.graphql;

import lombok.RequiredArgsConstructor;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class HealthGraphqlController {
    @QueryMapping
    public String health() {
        return "ok";
    }
}
