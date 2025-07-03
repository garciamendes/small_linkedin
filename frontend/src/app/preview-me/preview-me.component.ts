import { Component } from "@angular/core";
import { LucideAngularModule, ShieldCheck } from "lucide-angular";

@Component({
  selector: "app-preview-me",
  imports: [LucideAngularModule],
  templateUrl: "./preview-me.component.html",
  styleUrl: "./preview-me.component.css",
})
export class PreviewMeComponent {
  readonly ShieldCheck = ShieldCheck;

  previewThumb =
    "https://media.licdn.com/dms/image/v2/C4D16AQEnWRJrYbZ0HA/profile-displaybackgroundimage-shrink_350_1400/profile-displaybackgroundimage-shrink_350_1400/0/1605982758778?e=1756944000&v=beta&t=2QTs6-NYW34ens74jMVhGds-1bO7vC2kFyMNoHrFh_s";
  avatar =
    "https://media.licdn.com/dms/image/v2/D4D03AQEy8FODnawn6g/profile-displayphoto-shrink_200_200/B4DZWBLhbzG4AY-/0/1741629036794?e=1756944000&v=beta&t=ttoZ1XWIGeKIxopE1m5qRiWJi1AERi9VFYtRPtauVAM";
}
