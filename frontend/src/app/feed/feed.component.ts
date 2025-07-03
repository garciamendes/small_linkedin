import { Component } from "@angular/core";
import { NewPublicationComponent } from "../new-publication/new-publication.component";

@Component({
  selector: "app-feed",
  imports: [NewPublicationComponent],
  templateUrl: "./feed.component.html",
  styleUrl: "./feed.component.css",
})
export class FeedComponent {}
