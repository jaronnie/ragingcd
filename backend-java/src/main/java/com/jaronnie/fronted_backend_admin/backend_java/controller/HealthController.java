package com.jaronnie.fronted_backend_admin.backend_java.controller;

import com.jaronnie.fronted_backend_admin.backend_java.util.R;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/health")
public class HealthController {
    @GetMapping("")
    public R<String> health() {
        return R.ok("", "ok");
    }
}
