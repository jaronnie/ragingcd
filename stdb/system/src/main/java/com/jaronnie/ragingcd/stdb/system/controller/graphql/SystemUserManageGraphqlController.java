package com.jaronnie.ragingcd.stdb.system.controller.graphql;

import com.jaronnie.ragingcd.stdb.system.domain.query.PageQuery;
import com.jaronnie.ragingcd.stdb.system.domain.query.SearchUserQuery;
import com.jaronnie.ragingcd.stdb.system.domain.vo.TableDataInfo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.UserVo;
import com.jaronnie.ragingcd.stdb.system.service.IUserService;
import lombok.RequiredArgsConstructor;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
public class SystemUserManageGraphqlController {
    private final IUserService iUserService;

    @QueryMapping
    public TableDataInfo<UserVo> getUsers() {
        return iUserService.queryPageList(new PageQuery(), new SearchUserQuery());
    }
}

