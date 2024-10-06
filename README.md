**Crunch: Your Data Compression Powerhouse**

**Introduction**

Crunch is a robust command-line tool written in Go that empowers you to achieve efficient data compression using the venerable Huffman coding algorithm. It streamlines the process, making it a breeze to compress your files while maintaining clarity and control.

**Features**

- **Lossless Compression:** Crunch preserves the integrity of your data, ensuring an exact reconstruction upon decompression.
- **Customizable Verbosity:** Control the level of detail displayed during compression with the `-v` flag, allowing you to tailor the output to your preferences.
- **Streamlined File Handling:** Crunch seamlessly handles both input and output files, making compression workflows effortless.

** Usage **

1. **Basic Usage:**
   ```bash
   crunch -i input.txt -o compressed.dat
   ```
   This command compresses `input.txt` and stores the compressed data in `compressed.dat`.

2. **Example with Verbose Mode**

```bash
crunch -i sample_data.txt -o encoded.dat -v
```

This example compresses `sample_data.txt` while providing informative progress messages about the compression process.

**Contributing**

We welcome contributions to Crunch! Feel free to fork the repository on GitHub (link coming soon) and submit pull requests for bug fixes, enhancements, or documentation improvements.

**License**

Crunch is distributed under the MIT License (see LICENSE file).

**Disclaimer**

While Crunch offers effective compression, some file types might not experience significant size reduction due to inherent redundancies within the data itself.

**Future Enhancements**

- Support for multiple compression algorithms (e.g., LZMA)
- Integration with popular archive management tools
- Decompression functionality (coming soon)

**I appreciate your interest in Crunch!**

