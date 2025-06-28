import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { PreviewMeComponent } from "../preview-me/preview-me.component";
import { HeaderComponent } from "../header/header.component";

@Component({
  selector: "app-root-layout",
  imports: [RouterOutlet, PreviewMeComponent, HeaderComponent],
  templateUrl: "./root-layout.component.html",
  styleUrl: "./root-layout.component.css",
})
export class RootLayoutComponent {}
