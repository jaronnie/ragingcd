package com.jaronnie.fronted_backend_admin.backend_java.domain.vo;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class LoginResponseVo {
    private String token;
}
