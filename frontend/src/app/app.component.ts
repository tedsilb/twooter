import { Component, ChangeDetectionStrategy } from "@angular/core";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  changeDetection: ChangeDetectionStrategy.OnPush,
  styleUrls: ["./app.component.sass"],
})
export class AppComponent {
  title = "twooter";
}
