package com.jaronnie.ragingcd.stdb.system.domain.bo;

import lombok.Data;

@Data
public class RegisterUserBo {
    String username;
    String password;
    String email;
    String verifyCode;
}
