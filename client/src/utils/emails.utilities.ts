import Fuse from 'fuse.js';

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
  

export  const HighlighWord = (text: string, term: string) => {
    const regex = new RegExp(term, 'gi')
    return text.replace(
      regex,
      '<span class="bg-yellow-100 dark:bg-royal_purple dark:text-slate-950  opacity-80 font-bold text-black rounded-md">' + term + '</span>'
    )
  }

export  function SearchWithFuse(words :string[], searchTerm: string) {
    const fuse = new Fuse(words, { threshold: 0.2 });
    const results = fuse.search(searchTerm);

    if (results.length > 0) {
        return results
            .filter((result) => result.item.startsWith(searchTerm.slice(0, 2)))
            .map((result) => result.item);
    } else {
        return [];
    }
}