import 'package:currency_calc/modules/conversion/app/widget/CurrencyCalcApp.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('Element existence smoke test', (WidgetTester tester) async {
    await tester.pumpWidget(CurrencyCalcApp());
    expect(find.text('Enter amount'), findsOneWidget);
    expect(find.byKey(Key('sourceAmount')), findsOneWidget);
  });

  testWidgets('Wrong amount validation smoke test', (WidgetTester tester) async {
    // Arrange
    await tester.pumpWidget(CurrencyCalcApp());
    final txtSourceAmount = find.byKey(Key('sourceAmount'));
    // Act
    await tester.enterText(txtSourceAmount, 'hello');
    await tester.pump();
    // Assert
    expect(find.text('Numeric amount expected'), findsOneWidget);

    // Act
    await tester.enterText(txtSourceAmount, '-10');
    await tester.pump();
    // Assert
    expect(find.text('Enter positive non-zero amount'), findsOneWidget);
  });
}
