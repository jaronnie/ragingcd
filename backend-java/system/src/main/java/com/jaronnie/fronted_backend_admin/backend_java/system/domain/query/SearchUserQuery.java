package com.jaronnie.fronted_backend_admin.backend_java.system.domain.query;

import lombok.Data;

@Data
public class SearchUserQuery {
    private String username;
    private String email;
}
