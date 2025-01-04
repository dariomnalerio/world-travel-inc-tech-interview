import { createValidator, combineValidators, predicates } from '../../helpers/validators';

describe('createValidator', () => {
  it('should return null if the predicate is satisfied', () => {
    const predicate = (value: string) => value.length > 3;
    const validator = createValidator(predicate, 'testField', 'Value must be longer than 3 characters');

    const result = validator('valid');
    expect(result).toBeNull();
  });

  it('should return a ValidationError if the predicate is not satisfied', () => {
    const predicate = (value: string) => value.length > 3;
    const field = 'testField';
    const message = 'Value must be longer than 3 characters';
    const validator = createValidator(predicate, field, message);

    const result = validator('no');
    expect(result).toEqual({ field, message, });
  });

  it('should work with different types', () => {
    const predicate = (value: number) => value > 10;
    const field = 'testField';
    const message = 'Value must be greater than 10';
    const validator = createValidator(predicate, field, message);

    const result = validator(5);
    expect(result).toEqual({ field, message, });

    const validResult = validator(15);
    expect(validResult).toBeNull();
  });
});

describe('combineValidators', () => {
  const field1 = 'testField1';
  const message1 = 'Value must be longer than 3 characters';
  const field2 = 'testField2';
  const message2 = 'Value must contain an uppercase letter';

  const validator1 = createValidator((value: string) => value.length > 3, field1, message1);
  const validator2 = createValidator((value: string) => /[A-Z]/.test(value), field2, message2);

  const combinedValidator = combineValidators(validator1, validator2);
  const validValue = 'Valid';
  /**
   * Fails the first validator
   */
  const invalidValue1 = 'no';
  /**
   * Fails the second validator
   */
  const invalidValue2 = 'valid';
  it('should return null if all validators pass', () => {
    const result = combinedValidator(validValue);
    expect(result).toBeNull();
  });

  it('should return the first ValidationError if a validator fails', () => {
    const result = combinedValidator(invalidValue1);
    expect(result).toEqual({ field: 'testField1', message: 'Value must be longer than 3 characters' });
  });

  it('should return the second ValidationError if the first validator passes but the second fails', () => {
    const result = combinedValidator(invalidValue2);
    expect(result).toEqual({ field: 'testField2', message: 'Value must contain an uppercase letter' });
  });
});

describe("predicates", () => {

  describe("required", () => {
    it("should validate required fields", () => {
      const isValid = predicates.required("test");
      expect(isValid).toBe(true);
    })
    it("should fail to validate empty string", () => {
      const isValid = predicates.required("");
      expect(isValid).toBe(false);
    });
    it("should fail to validate undefined value", () => {
      // @ts-expect-error - Testing invalid input
      const isValid = predicates.required(undefined);
      expect(isValid).toBe(false);
    })
  })

  describe("minLength", () => {
    it("should validate strings with length greater than the minimum", () => {
      const isValid = predicates.minLength(3)("test");
      expect(isValid).toBe(true);
    });
    it("should validate strings with equal length to the minimum", () => {
      const isValid = predicates.minLength(4)("test");
      expect(isValid).toBe(true);
    });
    it("should fail to validate strings with length less than the minimum", () => {
      const isValid = predicates.minLength(5)("test");
      expect(isValid).toBe(false);
    });
  })

  describe("maxLength", () => {
    it("should validate strings with length less than the maximum", () => {
      const isValid = predicates.maxLength(5)("test");
      expect(isValid).toBe(true);
    });
    it("should validate strings with equal length to the maximum", () => {
      const isValid = predicates.maxLength(4)("test");
      expect(isValid).toBe(true);
    });
    it("should fail to validate strings with length greater than the maximum", () => {
      const isValid = predicates.maxLength(3)("test");
      expect(isValid).toBe(false);
    });
  })

  describe("hasUpperCase", () => {
    it("should validate strings with uppercase letters", () => {
      const isValid = predicates.hasUpperCase("Test");
      expect(isValid).toBe(true);
    });
    it("should validate strings with more than one uppercase letter", () => {
      const isValid = predicates.hasUpperCase("TesT");
      expect(isValid).toBe(true);
    });
    it("should validate strings with only uppercase letters", () => {
      const isValid = predicates.hasUpperCase("TEST");
      expect(isValid).toBe(true);
    });
    it("should fail to validate strings without uppercase letters", () => {
      const isValid = predicates.hasUpperCase("test");
      expect(isValid).toBe(false);
    });
  });

  describe("hasLowerCase", () => {
    it("should validate strings with lowercase letters", () => {
      const isValid = predicates.hasLowerCase("test");
      expect(isValid).toBe(true);
    });
    it("should validate strings with more than one lowercase letter", () => {
      const isValid = predicates.hasLowerCase("tesT");
      expect(isValid).toBe(true);
    });
    it("should validate strings with only lowercase letters", () => {
      const isValid = predicates.hasLowerCase("test");
      expect(isValid).toBe(true);
    });
    it("should fail to validate strings without lowercase letters", () => {
      const isValid = predicates.hasLowerCase("TEST");
      expect(isValid).toBe(false);
    });
  });

  describe("hasDigit", () => {
    it("should validate strings with digits", () => {
      const isValid = predicates.hasDigit("test1");
      expect(isValid).toBe(true);
    });
    it("should validate strings with more than one digit", () => {
      const isValid = predicates.hasDigit("tesT123");
      expect(isValid).toBe(true);
    });
    it("should validate strings with only digits", () => {
      const isValid = predicates.hasDigit("123");
      expect(isValid).toBe(true);
    });
    it("should fail to validate strings without digits", () => {
      const isValid = predicates.hasDigit("TEST");
      expect(isValid).toBe(false);
    });
  })

  describe("hasSpecialChar", () => {
    it("should validate strings with special characters", () => {
      const isValid = predicates.hasSpecialChar("test!");
      expect(isValid).toBe(true);
    });
    it("should validate strings with more than one special character", () => {
      const isValid = predicates.hasSpecialChar("tesT!@#");
      expect(isValid).toBe(true);
    });
    it("should validate strings with only special characters", () => {
      const isValid = predicates.hasSpecialChar("!@#");
      expect(isValid).toBe(true);
    });
    it("should fail to validate strings without special characters", () => {
      const isValid = predicates.hasSpecialChar("TEST");
      expect(isValid).toBe(false);
    });
  })

  describe("isEmail", () => {
    it("should validate valid email addresses", () => {
      const isValid = predicates.isEmail("test@test.com");
      expect(isValid).toBe(true);
    })
    it("should fail to validate invalid email addresses", () => {
      const isValid = predicates.isEmail("test");
      expect(isValid).toBe(false);
    })
    it("should fail to validate email addresses without a domain", () => {
      const isValid = predicates.isEmail("test@");
      expect(isValid).toBe(false);
    })
    it("should fail to validate email addresses without a username", () => {
      const isValid = predicates.isEmail("@test.com");
      expect(isValid).toBe(false);
    })
    it("should fail to validate email addresses without an extension", () => {
      const isValid = predicates.isEmail("test@.com");
      expect(isValid).toBe(false);
    })

  })
})