package com.jaronnie.ragingcd.stdb.system.controller.graphql;

import lombok.RequiredArgsConstructor;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class SystemHealthGraphqlController {
    @QueryMapping
    public String health() {
        return "ok";
    }
}
