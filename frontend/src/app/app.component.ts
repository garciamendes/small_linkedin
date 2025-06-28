import { Component, inject } from "@angular/core";
import { NavigationStart, Router, RouterOutlet } from "@angular/router";
import { NgxSonnerToaster } from "ngx-sonner";
import { AuthService } from "./core/services/auth.service";
import { filter } from "rxjs";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";

@Component({
  selector: "app-root",
  imports: [RouterOutlet, NgxSonnerToaster],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.css",
})
export class AppComponent {
  private authService = inject(AuthService);
  private router = inject(Router);

  constructor() {
    this.router.events
      .pipe(
        filter((e) => e instanceof NavigationStart),
        takeUntilDestroyed()
      )
      .subscribe((event: any) => {
        const url = event.url;
        const isLoggedIn = this.authService.getToken();

        if (url === "/login" && isLoggedIn) {
          this.router.navigate(["/feed"]);
        }

        if (url !== "/login" && !isLoggedIn) {
          this.router.navigate(["/login"]);
        }
      });
  }
}
