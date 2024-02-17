import {Component, OnInit} from '@angular/core';
import {user} from "../../models/users";
import {UserService} from "../../services/user.service";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.scss']
})
export class UsersComponent implements OnInit {
  displayedColumns: string[] = ['ID', 'UserID', 'CreatedAt', 'UpdatedAt', 'DeletedAt', 'actions'];
  dataSource: Array<user> = []

  newUserForm: FormGroup;
  isFormVisible = false;
  constructor(private userService: UserService, private fb: FormBuilder) {
    this.newUserForm = this.fb.group({
      userId: [null, Validators.required],
    });
  }
  addUser() {
    const userID = this.newUserForm?.value?.userId
    this.createUser(+userID)
    this.closeForm()
  }

  createUser(id: number) {
    this.userService.createUser(id).subscribe((res: { message: string, user: user}) => {
      const {user} = res
      const isHaveInArray = this.dataSource.findIndex(u => u.ID === user.ID);
      if (isHaveInArray !== -1) {
        this.dataSource = this.dataSource.map(u => (u.ID === user.ID ? user : u));
      } else {
        this.dataSource = [...this.dataSource, user];
      }
    })
  }
  deleteUser(id: number) {
    this.userService.deleteUser(id).subscribe((res: { message: string, user: user}) => {
      const {user} = res
      this.dataSource = this.dataSource.map(u => (u.ID === user.ID ? user : u));
    })
  }

  restoreUser(id:number) {
    this.createUser(id)
  }

  ngOnInit() {
    this.userService.getUsers().subscribe(({users}: {users: Array<user>} ) => {
      this.dataSource = users;
    });
  }

  openForm() {
    this.isFormVisible = true;
  }

  closeForm() {
    this.isFormVisible = false;
    this.newUserForm.reset();
  }
}
