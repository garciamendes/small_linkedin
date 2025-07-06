import { Routes } from "@angular/router";
import { NotFoundComponent } from "./not-found/not-found.component";
import { FeedComponent } from "./feed/feed.component";
import { authGuard } from "./core/guards/auth.guard";
import { RootLayoutComponent } from "./root-layout/root-layout.component";
import { NetworkComponent } from "./network/network.component";
import { AuthComponent } from "./auth/auth.component";

export const routes: Routes = [
  {
    path: "",
    redirectTo: "/feed",
    pathMatch: "full",
  },
  {
    path: "login",
    component: AuthComponent,
  },
  {
    path: "",
    component: RootLayoutComponent,
    canActivate: [authGuard],
    children: [
      {
        path: "feed",
        component: FeedComponent,
      },
      {
        path: "network",
        component: NetworkComponent,
      },
    ],
  },
  { path: "**", component: NotFoundComponent },
];
