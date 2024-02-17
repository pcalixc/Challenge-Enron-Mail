export function  ConvertDateFormat(originalDate: string) {
    // Create a Date object with the original date string
    const date = new Date(originalDate);
  
    // Array with the full names of the months
    const months = [
      'January',
      'February',
      'March',
      'April',
      'May',
      'June',
      'July',
      'August',
      'September',
      'October',
      'November',
      'December'
    ];
  
    // Get the full name of the month
    const monthName = months[date.getMonth()];
  
    // Get the day and year
    const day = date.getDate();
    const year = date.getFullYear();
  
    // Format the date in the "October 9, 2020" format
    const formattedDate = monthName + ' ' + day + ', ' + year;
  
    return formattedDate;
  }
  
export function SeparateEmailsByCommas(string: string) {
    return string.split(',').filter(function (substring) {
      return substring.trim() !== ''; // Exclude empty strings after trimming whitespace
    });
  }
  

export  const HighlighWord = (text: string, word: string) => {
    const regex = new RegExp(word, 'gi')
    return text.replace(
      regex,
      '<span class="bg-yellow-200 opacity-80 font-bold text-black rounded-full">' + word + '</span>'
    )
  }