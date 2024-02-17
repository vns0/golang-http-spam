import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-skeleton',
  template: `<div class="skeleton" [style.width]="width" [style.height]="height"></div>`,
  styleUrls: ['skeleton.component.scss'],
})
export class SkeletonComponent {
  @Input() width = '125px';
  @Input() height = '120px';
}
