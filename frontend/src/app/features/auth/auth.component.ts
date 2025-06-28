import { Component, inject } from "@angular/core";
import {
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from "@angular/forms";
import { AuthService } from "../../core/services/auth.service";
import { toast } from "ngx-sonner";
import { Router } from "@angular/router";

interface ILoginProps {
  email: string;
  password: string;
}

@Component({
  selector: "app-auth",
  imports: [FormsModule, ReactiveFormsModule],
  templateUrl: "./auth.component.html",
  styleUrl: "./auth.component.css",
})
export class AuthComponent {
  loginForm: FormGroup;
  private authService = inject(AuthService);
  private router = inject(Router);

  constructor(private fb: FormBuilder) {
    this.loginForm = this.fb.group({
      email: ["", [Validators.required, Validators.email]],
      password: ["", [Validators.required]],
    });
  }

  get email() {
    return this.loginForm.get("email");
  }

  get password() {
    return this.loginForm.get("password");
  }

  onSubmit() {
    if (!this.loginForm.valid) {
      toast.error("Todos os Campos são obrigatórios");
      return;
    }

    const data = this.loginForm.value;
    this.authService.login(data).subscribe({
      next: () => this.router.navigate(["/feed"]),
      error: (err) => toast.error(err.error.message),
    });
  }
}
