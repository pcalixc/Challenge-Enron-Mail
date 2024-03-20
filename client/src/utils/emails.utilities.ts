import Fuse from 'fuse.js';

export const ConvertDateFormat =(originalDate: string): string => {
  const year = parseInt(originalDate.substring(0, 4));
  const month = parseInt(originalDate.substring(5, 7));
  const day = parseInt(originalDate.substring(8, 10));
  const hour = parseInt(originalDate.substring(11, 13));
  const minute = parseInt(originalDate.substring(14, 16));
  const second = parseInt(originalDate.substring(17, 19));

  const fecha = new Date(year, month - 1, day, hour, minute, second);

  const formatOptions: Intl.DateTimeFormatOptions = { month: 'long', day: 'numeric', year: 'numeric', hour: '2-digit', minute: '2-digit' };
  const formatedDate = fecha.toLocaleDateString('en-US', formatOptions);
  return formatedDate;
}

export const SeparateEmailsByCommas=(string: string) =>{
    return string.split(',').filter(function (substring) {
      return substring.trim() !== ''; // Exclude empty strings after trimming whitespace
    });
  }
  
export const HighlighWord = (text: string, term: string) => {
    const regex = new RegExp(term, 'gi')
    return text.replace(
      regex,
      '<span class="bg-yellow-100 dark:text-slate-900  opacity-80 font-bold text-black rounded-md">' + term + '</span>'
    )
  }

export const SearchWithFuse=(words :string[], searchTerm: string) =>{
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