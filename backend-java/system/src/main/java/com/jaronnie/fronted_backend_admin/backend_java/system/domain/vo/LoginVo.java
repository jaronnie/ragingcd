package com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class LoginVo {
    private String token;
}