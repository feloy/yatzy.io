import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'yatzyScore'
})
export class YatzyScorePipe implements PipeTransform {

  transform(value: number, args?: any): any {
    if (value == null) {
      return '';
    }
    const l = 4 - value.toString().length;
    return Array(l + 1).join('0');
  }

}
