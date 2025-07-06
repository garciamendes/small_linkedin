import { Component, computed, signal } from "@angular/core";
import { Image, LucideAngularModule, X } from "lucide-angular";

@Component({
  selector: "app-new-publication",
  imports: [LucideAngularModule],
  templateUrl: "./new-publication.component.html",
  styleUrl: "./new-publication.component.css",
})
export class NewPublicationComponent {
  icons = {
    Close: X,
    Image: Image,
  } as const;

  avatar =
    "https://media.licdn.com/dms/image/v2/D4D03AQEy8FODnawn6g/profile-displayphoto-shrink_200_200/B4DZWBLhbzG4AY-/0/1741629036794?e=1756944000&v=beta&t=ttoZ1XWIGeKIxopE1m5qRiWJi1AERi9VFYtRPtauVAM";

  isOpenModalNewPublication = signal(false);
  publicationContent = signal("");
  isDisabledButton = computed(
    () => !this.publicationContent().trim().length && !this.file()
  );

  file = signal<File | null>(null);
  filePreview = signal<string | null>(null);
  fileProgress = signal(0);
  fileLoading = signal(false);

  closeModalNewPublication() {
    this.isOpenModalNewPublication.set(false);
    this.publicationContent.set("");
    this.handleClearFile();
  }

  openModalNewPublication() {
    this.isOpenModalNewPublication.set(true);
  }

  handlerSubmit() {
    console.log(this.publicationContent(), this.file());
  }

  handleTriggerFile() {
    const input = document.getElementById("fileUpload") as HTMLInputElement;
    input.click();
  }

  handleFileChange(event: Event) {
    const target = event.target as HTMLInputElement;
    const file = target.files?.[0];
    if (!file) return;

    if (!file.type.startsWith("image/")) return;

    const reader = new FileReader();

    this.fileLoading.set(true);
    this.fileProgress.set(0);

    let simulatedProgress = 0;

    const interval = setInterval(() => {
      simulatedProgress += 1;
      if (simulatedProgress <= 90 && this.fileLoading()) {
        this.fileProgress.set(simulatedProgress);
      } else {
        clearInterval(interval);
      }
    }, 60);

    reader.onload = () => {
      clearInterval(interval);

      const finalInterval = setInterval(() => {
        simulatedProgress += 2;
        if (simulatedProgress < 100) {
          this.fileProgress.set(simulatedProgress);
        } else {
          clearInterval(finalInterval);
          this.fileProgress.set(100);

          setTimeout(() => {
            this.filePreview.set(reader.result as string);
            this.file.set(file);
            this.fileLoading.set(false);
            this.fileProgress.set(0);
          }, 300);
        }
      }, 30);
    };

    reader.onerror = () => {
      clearInterval(interval);
      this.fileLoading.set(false);
      this.fileProgress.set(0);
    };

    reader.readAsDataURL(file);
  }

  handleClearFile() {
    this.file.set(null);
    this.filePreview.set(null);
    this.fileProgress.set(0);
    this.fileLoading.set(false);

    const input = document.getElementById("fileUpload") as HTMLInputElement;
    if (input) {
      input.value = "";
    }
  }
}
