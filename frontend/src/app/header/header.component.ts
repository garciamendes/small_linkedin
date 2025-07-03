import { Component, inject } from "@angular/core";
import { Router, RouterModule } from "@angular/router";
import {
  House,
  Users,
  LucideAngularModule,
  LucideIconData,
} from "lucide-angular";

export interface INavigation {
  path: string;
  title: string;
  icon: LucideIconData;
}

@Component({
  selector: "app-header",
  imports: [LucideAngularModule, RouterModule],
  templateUrl: "./header.component.html",
  styleUrl: "./header.component.css",
})
export class HeaderComponent {
  readonly House = House;
  readonly Users = Users;

  private route = inject(Router);

  navigation: INavigation[] = [
    {
      path: "/feed",
      icon: this.House,
      title: "Início",
    },
    {
      path: "/network",
      icon: this.Users,
      title: "Rede",
    },
  ];

  isActiveRouter(path: string): boolean {
    return this.route.url === path;
  }
}
