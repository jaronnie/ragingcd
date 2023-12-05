package com.jaronnie.fronted_backend_admin.backend_java.controller;

import com.jaronnie.fronted_backend_admin.backend_java.util.R;
import lombok.RequiredArgsConstructor;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/health")
public class HealthController {
    @QueryMapping
    public R<String> healthByGraphql() {
        return R.ok("", "ok");
    }

    @GetMapping("")
    public R<String> health() {
        return R.ok("", "ok");
    }
}
