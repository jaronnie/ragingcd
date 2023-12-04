import type { RouteRecordRaw } from "vue-router";

export interface UserState {
  token: string | null;
  username: string;
  avatar: string;
  menuRoutes: RouteRecordRaw[];
}

export interface LayoutState {
  fold: boolean;
  refresh: boolean;
}
