import { Component } from "@angular/core";
import { House, LucideAngularModule } from "lucide-angular";

@Component({
  selector: "app-header",
  imports: [LucideAngularModule],
  templateUrl: "./header.component.html",
  styleUrl: "./header.component.css",
})
export class HeaderComponent {
  readonly House = House;
}
