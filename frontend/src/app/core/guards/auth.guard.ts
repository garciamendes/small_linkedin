import { inject } from "@angular/core";
import { CanActivateFn, Router } from "@angular/router";
import { AuthService } from "../services/auth.service";

export const authGuard: CanActivateFn = (route, state) => {
  const authToken = inject(AuthService).getToken();
  const router = inject(Router);

  if (!authToken) {
    router.navigate(["/login"]);
    return false;
  }

  return true;
};
