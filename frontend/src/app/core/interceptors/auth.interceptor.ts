import { HttpInterceptorFn } from "@angular/common/http";
import { inject } from "@angular/core";
import { AuthService } from "../services/auth.service";

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const token = inject(AuthService).getToken();

  if (token) {
    const clonned = req.clone({
      headers: req.headers.set("Authorization", `Bearer ${token}`),
    });

    return next(clonned);
  }

  return next(req);
};
