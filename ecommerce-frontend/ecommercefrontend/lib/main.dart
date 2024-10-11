import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Data Table Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Flutter Data Table Page'),
    );
  }
}

class Product {
  final int id;
  final String name;
  final double price;

  Product({required this.id, required this.name, required this.price});

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['id'],
      name: json['name'],
      price: (json['price'] is String)
          ? double.tryParse(json['price']) ??
              0.0 // Convert string to double if necessary
          : json['price'].toDouble(), // Ensure price is double
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  Future<List<Product>>? futureProducts;

  Future<List<Product>> fetchProducts() async {
    final response =
        await http.get(Uri.parse('http://192.168.1.108:8082/products'));

    if (response.statusCode == 200) {
      List jsonResponse = json.decode(response.body);
      return jsonResponse.map((product) => Product.fromJson(product)).toList();
    } else {
      throw Exception('Failed to load products');
    }
  }

  void refreshData() {
    setState(() {
      futureProducts = fetchProducts(); // Re-fetch data
    });
  }

  @override
  void initState() {
    super.initState();
    refreshData(); // Initialize in initState
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: refreshData, // Call refreshData on button press
          ),
        ],
      ),
      body: Center(
        child: FutureBuilder<List<Product>>(
          future: futureProducts,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const CircularProgressIndicator(); // Show loading spinner while waiting
            } else if (snapshot.hasError) {
              return Text("${snapshot.error}"); // Show error if any
            } else if (snapshot.hasData) {
              return SingleChildScrollView(
                scrollDirection: Axis.horizontal,
                child: DataTable(
                  columns: const [
                    DataColumn(label: Text('ID')),
                    DataColumn(label: Text('Name')),
                    DataColumn(label: Text('Price')),
                  ],
                  rows: snapshot.data!
                      .map((product) => DataRow(cells: [
                            DataCell(Text(product.id.toString())),
                            DataCell(Text(product.name)),
                            DataCell(Text(product.price.toStringAsFixed(
                                2))), // Format to 2 decimal places
                          ]))
                      .toList(),
                ),
              );
            } else {
              return const Text('No data available');
            }
          },
        ),
      ),
    );
  }
}
