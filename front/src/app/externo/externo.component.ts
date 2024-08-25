import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatToolbarModule} from '@angular/material/toolbar';

@Component({
  selector: 'app-externo',
  standalone: true,
  imports: [ RouterOutlet, MatToolbarModule, MatButtonModule, MatIconModule ],
  templateUrl: './externo.component.html',
  styleUrl: './externo.component.css'
})
export class ExternoComponent {

}
