export function nucleotideCounts(input: string) {
  let result = { A: 0, C: 0, G: 0, T: 0 };
  [...input].forEach((c, i) => {
    switch (c) {
      case "A":
        result.A++;
        break;
      case "C":
        result.C++;
        break;
      case "G":
        result.G++;
        break;
      case "T":
        result.T++;
        break;
      default:
        throw new Error('Invalid nucleotide in strand');
    }
  });
  return result;
}
