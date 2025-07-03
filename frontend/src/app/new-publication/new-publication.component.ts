import {
  Component,
  effect,
  ElementRef,
  signal,
  ViewChild,
} from "@angular/core";
import { Image, LucideAngularModule, X } from "lucide-angular";

@Component({
  selector: "app-new-publication",
  imports: [LucideAngularModule],
  templateUrl: "./new-publication.component.html",
  styleUrl: "./new-publication.component.css",
})
export class NewPublicationComponent {
  readonly Close = X;
  readonly Image = Image;

  avatar =
    "https://media.licdn.com/dms/image/v2/D4D03AQEy8FODnawn6g/profile-displayphoto-shrink_200_200/B4DZWBLhbzG4AY-/0/1741629036794?e=1756944000&v=beta&t=ttoZ1XWIGeKIxopE1m5qRiWJi1AERi9VFYtRPtauVAM";

  isOpenModalNewPublication = signal(true);
  isDisabledButton = true;

  closeModalNewPublication() {
    this.isOpenModalNewPublication.set(false);
  }

  openModalNewPublication() {
    this.isOpenModalNewPublication.set(true);
  }
}
