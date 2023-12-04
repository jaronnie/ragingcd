import type { RouteRecordRaw } from "vue-router";

export interface UserState {
  token: string | null;
  username: string;
  menuRoutes: RouteRecordRaw[];
}

export interface LayoutState {
  fold: boolean;
  refresh: boolean;
}
