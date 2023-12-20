package com.jaronnie.fronted_backend_admin.backend_java.system.controller;

import com.jaronnie.fronted_backend_admin.backend_java.common.util.R;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

@Api(tags = "服务运行状态管理")
@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/system/health")
public class SystemHealthController {
    @ApiOperation(value = "获取服务健康状态")
    @GetMapping("")
    public R<String> health() {
        return R.ok("", "ok");
    }
}
