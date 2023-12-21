package com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo;

import lombok.Data;

@Data
public class RegisterUserBo {
    String username;
    String password;
    String email;
}
