package com.jaronnie.ragingcd.stdb.system.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.jaronnie.ragingcd.stdb.system.domain.po.UserPo;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface UserMapper extends BaseMapper<UserPo> {

}

