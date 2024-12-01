import { Component } from "@angular/core";
import { NgSwitch, NgSwitchCase } from "@angular/common";

@Component({
  selector: "app-root",
  imports: [
    NgSwitch,
    NgSwitchCase,
  ],
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.sass"],
})
export class AppComponent {
  title = "twooter";
}
