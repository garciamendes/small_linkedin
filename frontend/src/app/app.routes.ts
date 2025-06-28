import { Routes } from "@angular/router";
import { AuthComponent } from "./features/auth/auth.component";
import { NotFoundComponent } from "./not-found/not-found.component";
import { FeedComponent } from "./feed/feed.component";
import { authGuard } from "./core/guards/auth.guard";

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
    path: "feed",
    component: FeedComponent,
    canActivate: [authGuard],
  },
  { path: "**", component: NotFoundComponent },
];
