Nama: Muhammad Ridhan Khoirullah
ID Camp: BE9259656

Pertama, didefinisikan sebuah antarmuka Produk yang memiliki satu metode TampilkanInfo() yang akan memberikan informasi tentang produk tertentu. Kemudian, ada dua struktur yang mengimplementasikan antarmuka Produk:
Keripik dan Kerupuk. Masing-masing memiliki properti Nama yang merepresentasikan jenis produknya.

Selanjutnya, ada antarmuka Pabrik yang memiliki satu metode BuatProduk() yang mengembalikan sebuah produk. Dua pabrik konkret diimplementasikan:
PabrikKeripik dan PabrikKerupuk. Setiap pabrik memiliki metode BuatProduk() yang mengembalikan occurrence baru dari jenis produk yang sesuai.

Di dalam fungsi fundamental(), dilakukan penggunaan pabrik untuk membuat occasion produk dari jenis keripik dan kerupuk. Pertama, sebuah occurrence dari PabrikKeripik dibuat dan digunakan untuk membuat produk keripik. Kemudian, hal yang sama dilakukan dengan PabrikKerupuk untuk membuat produk kerupuk. Setelah itu, informasi tentang produk-produk yang telah dibuat ditampilkan menggunakan metode TampilkanInfo() yang telah diimplementasikan pada masing-masing produk.

Melalui desain ini, aplikasi dapat secara dinamis membuat berbagai jenis produk tanpa harus terikat langsung ke implementasi spesifiknya, sehingga meningkatkan fleksibilitas dan memudahkan pengelolaan kode. 