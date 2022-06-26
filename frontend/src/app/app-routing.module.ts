import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { JobsComponent } from './components/jobs/jobs.component';
import { ProfilesComponent } from './components/profiles/profiles.component';

const routes: Routes = [
  { path: 'jobs', component: JobsComponent },
  { path: 'profiles' , component: ProfilesComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
