import {Component, OnInit} from '@angular/core';
import {DashboardService} from "../../services/dashboard.service";
import {statsResponse} from "../../models/dashboard";

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  attackCount: number = 0;
  userCount: number = 0;
  numberOfAttacks: number = 0;
  isLoadingStats = false

  constructor(private dshboardService: DashboardService) {}


  startAttack() {
    console.log('click')
  }

  loadStats() {
    this.isLoadingStats = true;

    this.dshboardService.getStats().subscribe(
      (data: statsResponse) => {
        console.log(data)
        this.attackCount = data.countAttack;
        this.userCount = data.countUsers;
      },
      error => console.error(error),
      () => {
        // Устанавливаем isLoadingStats в false после завершения запроса
        this.isLoadingStats = false;
      }
    );
  }

  ngOnInit(): void {
    this.loadStats()
  }
}
